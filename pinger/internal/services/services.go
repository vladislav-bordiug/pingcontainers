package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"pinger/internal/models"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	backendURL string
}

func NewService(backendURL string) *Service {
	return &Service{backendURL: backendURL}
}

func (s *Service) RunPingCycle() {
	out, err := exec.Command("docker", "ps", "-q").Output()
	if err != nil {
		log.Printf("Ошибка получения списка контейнеров: %v", err)
		return
	}

	containerIDs := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, id := range containerIDs {
		if id == "" {
			continue
		}

		ip, err := getContainerIP(id)
		if err != nil {
			log.Printf("Ошибка получения IP для контейнера %s: %v", id, err)
			continue
		}

		var lastSuccess time.Time
		pingTime, err := pingIP(ip)
		if err != nil {
			log.Printf("Ping %s не успешен: %v", ip, err)
		} else {
			lastSuccess = time.Now()
		}

		ps := models.PingStatus{
			IP:          ip,
			PingTime:    pingTime,
			LastSuccess: lastSuccess,
		}

		sendPingStatus(s.backendURL, ps)
	}
}

func getContainerIP(containerID string) (string, error) {
	out, err := exec.Command("docker", "inspect", "-f", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", containerID).Output()
	if err != nil {
		return "", err
	}
	ip := strings.TrimSpace(string(out))
	if ip == "" {
		return "", nil
	}
	return ip, nil
}

func pingIP(ip string) (float64, error) {
	out, err := exec.Command("ping", "-c", "1", "-W", "1", ip).CombinedOutput()
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`time=([\d\.]+) ms`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) < 2 {
		return 0, err
	}
	return strconv.ParseFloat(matches[1], 64)
}

func sendPingStatus(backendURL string, ps models.PingStatus) {
	data, err := json.Marshal(ps)
	if err != nil {
		log.Printf("Ошибка маршалинге JSON: %v", err)
		return
	}
	url := backendURL + "/api/ping"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Ошибка отправки данных в %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Неверный статус ответа: %s", resp.Status)
	}
}

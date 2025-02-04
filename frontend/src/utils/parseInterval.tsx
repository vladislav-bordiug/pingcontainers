export function parseInterval(intervalStr: string): number {
    console.log(intervalStr);
    if (intervalStr.endsWith("s")) {
      return parseInt(intervalStr) * 1000;
    }
    if (intervalStr.endsWith("ms")) {
      return parseInt(intervalStr);
    }
    return parseInt(intervalStr);
  }
  
export const intervaltime = parseInterval(import.meta.env.VITE_PINGER_INTERVAL_SECONDS || '30s');
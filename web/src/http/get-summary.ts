interface SummaryResponse {
  completed: number;
  total: number;
  goals_per_day: GoalsPerDay;
}

type GoalsPerDay = Record<string, GoalsPerDayData[]>;

interface GoalsPerDayData {
  id: string;
  title: string;
  completed_at: string;
}
export async function getSummary(): Promise<SummaryResponse> {
  return await fetch(`${import.meta.env.VITE_API_URL}/summary`)
    .then((res) => res.json());
}

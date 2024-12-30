export interface PendingGoalsResponse {
  id: string;
  title: string;
  desired_weekly_frequency: number;
  created_at: string;
  completion_count: number;
}

export async function getPendingGoals(): Promise<PendingGoalsResponse[]> {
  return await fetch(`${import.meta.env.VITE_API_URL}/pending-goals`)
    .then((res) =>
      res.json(),
    );
}

interface CreateGoalRequest {
  title: string;
  desiredWeeklyFrequency: number;
}

export interface CreateGoalResponse {
  id: string;
  title: string;
  desired_weekly_frequency: number;
  created_at: string;
}

export async function createGoal({
  title,
  desiredWeeklyFrequency,
}: CreateGoalRequest): Promise<CreateGoalResponse> {
  return await fetch(`${import.meta.env.VITE_API_URL}/goal`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      title,
      desired_weekly_frequency: desiredWeeklyFrequency,
    }),
  }).then((res) => res.json());
}

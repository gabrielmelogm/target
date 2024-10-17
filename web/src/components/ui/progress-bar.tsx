import * as ProgressPrimitive from "@radix-ui/react-progress";

export function Progress(props: ProgressPrimitive.ProgressProps) {
  return (
    <ProgressPrimitive.Progress
      className="bg-zinc-900 rounded-full h-2"
      {...props}
    />
  );
}

export function ProgressIndicator(
  props: ProgressPrimitive.ProgressIndicatorProps,
) {
  return (
    <ProgressPrimitive.Indicator
      className="bg-gradient-to-r from-pink-500 to-violet-500 w-1/2 h-2 rounded-full"
      {...props}
    />
  );
}
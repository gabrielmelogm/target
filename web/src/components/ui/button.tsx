import { ComponentProps, forwardRef } from "react";
import { tv, VariantProps } from "tailwind-variants";

const button = tv({
  base: "flex items-center justify-center gap-2 rounded-lg text-sm font-medium tracking-tight outline-none ring-offset-2 ring-offset-black",
  variants: {
    variant: {
      primary:
        "bg-violet-500 text-violet-50 hover:bg-violet-600 ring-violet-500",
      secondary: "bg-zinc-900 text-zinc-400 hover:bg-zinc-800 bg-zinc-900",
    },
    size: {
      default: "px-4 py-2.5",
      sm: "px-3 py-1.5",
    },
  },
  defaultVariants: {
    variant: "primary",
    size: "default",
  },
});

type ButtonProps = ComponentProps<"button"> & VariantProps<typeof button>;

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, ...props }, ref) => {
    return (
      <button
        className={button({ variant, size, className })}
        ref={ref}
        {...props}
      />
    );
  },
);

Button.displayName = "Button";

export { Button };

import * as React from "react";
import { cn } from "@/lib/utils";

export interface BadgeProps extends React.HTMLAttributes<HTMLSpanElement> {
  variant?: "success" | "warning" | "danger" | "info" | "default";
}

export function Badge({ className, variant = "default", ...props }: BadgeProps) {
  const baseStyles = "inline-flex items-center px-2.5 py-0.5 rounded text-xs font-semibold transition-colors";
  
  const variants = {
    default: "bg-zinc-100 text-zinc-800 dark:bg-zinc-800 dark:text-zinc-300",
    success: "bg-emerald-100 text-emerald-800 dark:bg-emerald-950/40 dark:text-emerald-400",
    warning: "bg-amber-100 text-amber-800 dark:bg-amber-950/40 dark:text-amber-400",
    danger: "bg-rose-100 text-rose-800 dark:bg-rose-950/40 dark:text-rose-400",
    info: "bg-blue-100 text-blue-800 dark:bg-blue-950/40 dark:text-blue-400",
  };

  return (
    <span
      className={cn(baseStyles, variants[variant], className)}
      {...props}
    />
  );
}

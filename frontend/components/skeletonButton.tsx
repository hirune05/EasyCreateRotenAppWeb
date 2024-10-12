import { Button } from "@/components/ui/button";
import React, { ReactNode } from "react";

interface SkeletonButtonProps {
  children: ReactNode;  
}

export default function SkeletonButton({ children }: SkeletonButtonProps) {
  return (
    <Button size="icon" variant="ghost" className="rounded-full">
      {children}
    </Button>
  );
}

import type { PropsWithChildren } from "react";

export default function NestedLayout(props: PropsWithChildren) {
  return (
    <div>
      <p>This layout is nested </p>
      <div>{props.children}</div>
    </div>
  );
}

import type { PropsWithChildren } from "react";

export default function AppLayout(props: PropsWithChildren) {
  return (
    <>
      <header>Header</header>
      <main>{props.children}</main>
      <footer>Footer</footer>
    </>
  );
}

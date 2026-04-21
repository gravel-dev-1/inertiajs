import { createInertiaApp, type ResolvedComponent } from "@inertiajs/react";
import "../css/app.css";

createInertiaApp({
  title: (title) => `Gravel — React`,
  resolve: async (name) => {
    const pages = import.meta.glob<ResolvedComponent>("./pages/**/*.tsx");
    const page = pages[`./pages/${name}.tsx`];
    if (!page) throw new Error(`Page not found: ${name}`);
    return await page();
  },
  strictMode: true,
  progress: {
    color: "#4B5563",
  },
});

import { createInertiaApp, type ResolvedComponent } from "@inertiajs/svelte";
import "../css/app.css";

createInertiaApp({
  title: (title) => `${title} — Svelte`,
  resolve: async (name) => {
    const pages = import.meta.glob<ResolvedComponent>("./pages/**/*.svelte");
    const page = pages[`./pages/${name}.svelte`];
    if (!page) throw new Error(`Page not found: ${name}`);
    return await page();
  },
});

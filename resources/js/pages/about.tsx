import AppLayout from "@/layouts/app-layout";
import NestedLayout from "@/layouts/nested-layout";
import { Head } from "@inertiajs/react";

function About(props: unknown) {
  return (
    <>
      <Head title="About">
        <meta name="description" content="Golang @inertia/react Demo" />
        <meta name="author" content="John Doe" />
      </Head>
      <div>About</div>
      <pre>{JSON.stringify(props)}</pre>
    </>
  );
}

About.layout = [AppLayout, NestedLayout];

export default About;

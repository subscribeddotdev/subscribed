import Head from "next/head";
import React from "react";

interface Props {
  title: string;
  isHome?: boolean;
  description?: string;
}

export function PageMeta({ title, isHome = false, description }: Props) {
  return (
    <Head>
      <title>{isHome ? title : `${title} - Subscribed`}</title>
      {description && <meta name="description" content={description} />}
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <link rel="icon" href="/favicon.ico" />
      <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png" />
      <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png" />
      <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png" />
      <link rel="manifest" href="/site.webmanifest"></link>
    </Head>
  );
}

import * as React from "react";
import { AppProps } from "next/app";
import { ThemeProvider } from "@mui/material";
import { CacheProvider } from "@emotion/react";
import CssBaseline from '@mui/material/CssBaseline';
import createCache from "@emotion/cache";
import theme from "../theme";
import Layout from "../components/layout/Layout";
import { UserProvider } from '@auth0/nextjs-auth0/client';

export const cache = createCache({ key: "css", prepend: true });

export default function MyApp(props: AppProps): JSX.Element {
  const { Component, pageProps } = props;

  React.useEffect(() => {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector("#jss-server-side");
    if (jssStyles) {
      jssStyles.parentElement?.removeChild(jssStyles);
    }
  }, []);

  return (
    <UserProvider>
    <CacheProvider value={cache}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </ThemeProvider>
    </CacheProvider>
  </UserProvider>
  );
}
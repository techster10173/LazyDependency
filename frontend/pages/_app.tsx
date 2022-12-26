import * as React from "react";
import { AppProps } from "next/app";
import { ThemeProvider } from "@mui/material";
import { CacheProvider, EmotionCache } from "@emotion/react";
import CssBaseline from '@mui/material/CssBaseline';
import theme from "../config/theme";
import Layout from "../components/layout/Layout";
import { UserProvider } from '@auth0/nextjs-auth0/client';
import createEmotionCache from "../config/emotioncache";

export const cache = createEmotionCache();

interface NewAppProps extends AppProps {
  emotionCache?: EmotionCache;
}

export default function MyApp(props: NewAppProps): JSX.Element {
  const { Component, pageProps, emotionCache = cache } = props;

  React.useEffect(() => {
    // Remove the server-side injected CSS.
    // const jssStyles = document.querySelector("#jss-server-side");
    // if (jssStyles) {
    //   jssStyles.parentElement?.removeChild(jssStyles);
    // }
  }, []);

  return (
    <UserProvider>
      <CacheProvider value={emotionCache}>
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
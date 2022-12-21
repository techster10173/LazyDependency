import {
    Typography,
    Container,
    Grid,
    Box,
    Divider,
    Button,
  } from "@material-ui/core";
  import { useRouter } from "next/router";
  import Link from "next/link";
  
  export default function Header(): JSX.Element {
    const router = useRouter();
  
    return (
      <Box
        sx={{
          backgroundColor: (theme) => theme.palette.background.default,
        }}
      >
        <Container maxWidth="md" sx={{ py: 1 }}>
          <Grid container alignItems="center">
            <Grid item xs={2}>
            <Link href="/" passHref>
              <Typography variant="body1" align="center" sx={{ fontWeight: 600 }}>
                Lazy Dependency
              </Typography>
              </Link>
            </Grid>
            <Grid container item xs={10} justifyContent="flex-end">
              <Link href="/page2" passHref>
                <Button
                  sx={{ mr: 2 }}
                  color={router.pathname === "/" ? "primary" : "secondary"}
                  

                >
                  Login
                </Button>
              </Link>
              <Link href="/page3" passHref>
                <Button
                  sx={{ mr: 2 }}
                  color={router.pathname === "/" ? "primary" : "secondary"}
                >
                  Register
                </Button>
              </Link>
            </Grid>
          </Grid>
        </Container>
        <Divider />
      </Box>
    );
  }
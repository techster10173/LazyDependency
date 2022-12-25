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
  import { useUser } from '@auth0/nextjs-auth0/client';
  import React, { useState } from 'react';
  
  
  export default function Header(): JSX.Element {
    const router = useRouter();
    const [isOpen, setIsOpen] = useState(false);
    const { user, isLoading } = useUser();
    const toggle = () => setIsOpen(!isOpen);

  
    return (
      <Box
       
      >
        <Container >
          <Grid container alignItems="center">
            <Grid item xs={2}>
            <Link href="/" passHref>
            <Container>
              <img src="/logoFull.png" width={150} height={60}  />
            </Container>
              </Link>
            </Grid>
            <Grid container item xs={10} justifyContent="flex-end">
            {!isLoading && !user && (
              
                <Button
                  href="/api/auth/login"
                  variant="contained"
                  >
                  Log in
                </Button>
              
            )}
             {user && (
             
                  <Button
                    href="/api/auth/logout"
                    variant="contained">
                    Log out
                  </Button>
               
            )}
            </Grid>
          </Grid>
        </Container>
        <Divider />
      </Box>
    );
  }
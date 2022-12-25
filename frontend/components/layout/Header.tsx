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
        <Container maxWidth="md" >
          <Grid container alignItems="center">
            <Grid item xs={2}>
            <Link href="/" passHref>
              <Typography variant="body1" align="center" >
                Lazy Dependency
              </Typography>
              </Link>
            </Grid>
            <Grid container item xs={10} justifyContent="flex-end">
            {!isLoading && !user && (
              
                <Button
                  href="/api/auth/login"
                  >
                  Log in
                </Button>
              
            )}
             {user && (
             
                  <Button
                    href="/api/auth/logout">
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
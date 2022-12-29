import React from "react";
import { Typography, Box, Grid, Container, Button } from "@mui/material";
import Image from "next/image";

export default function ContactSection(): JSX.Element {
  return (
    <Container maxWidth="md" style={{ marginTop: '-20px' }} >
        <Grid container alignItems="center">
          <Grid item xs={12} sm={6}>
            <Typography variant="h2">Contact Us</Typography>
            <Typography variant="body1" >
              Get in touch with the Lazy Dependency team!
            </Typography>
            <Button variant="contained" href="mailto:avnerlipszyc36@gmail.com">
              Get in Touch
            </Button>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Container>
              <Image src="/assets/contact.svg" width={400} height={400} alt={""}  />
            </Container>
          </Grid>
        </Grid>
     
    </Container>
  );
}


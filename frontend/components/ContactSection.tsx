import React from "react";
import { Typography, Box, Grid, Container, Button } from "@material-ui/core";
import Image from "next/image";

export default function ContactSection(): JSX.Element {
  return (
    <Container maxWidth="md" >
      <Box sx={{  borderRadius: 4, p: 2 }}>
        <Grid container alignItems="center" spacing={4}>
          <Grid item xs={12} sm={6}>
            <Typography variant="h2">Contact Us</Typography>
            <Typography variant="body1" >
              Get in touch with the Lazy Dependency team!
            </Typography>
            <Button variant="contained" href="avnerlipszyc36@gmail.com">
              Get in Touch
            </Button>
          </Grid>
          <Grid item xs={12} sm={6}>
            <Container>
              <img src="/assets/contact.svg" width={400} height={400}  />
            </Container>
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
}
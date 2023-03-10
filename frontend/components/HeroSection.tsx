import React from "react";
import { Typography, Box, Grid, Container } from "@mui/material";
import Image from 'next/image'

type Props = {
  title: string;
  subtitle: string;
  image: string;
  children?: React.ReactNodeArray;
};

export default function HeroSection({
  title,
  subtitle,
  image,
  children,
}: Props): JSX.Element {
  return (
    <Box >
      <Container maxWidth="md">
        <Grid container alignItems="center" >
          <Grid item xs={12} sm={6}>
            <Typography variant="h1">{title}</Typography>
            <Typography variant="body1">{subtitle}</Typography>
            {children}
          </Grid>
          <Grid item xs={12} sm={6}>
            <Container >
              <img 
              src={image} 
              width={400} 
              height={400} 
              />
            </Container>
          </Grid>
        </Grid>
      </Container>
    </Box>
  );
}
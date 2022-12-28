import React from "react";
import { Grid, Container } from "@mui/material";

interface Props {
  children: React.ReactNodeArray;
}

export default function FeatureBlocksContainer({
  children,
}: Props): JSX.Element {
  return (
    <Container  >
    <Grid container spacing={4}  marginLeft='25px' >
      {children}
    </Grid>
    </Container>
  );
}
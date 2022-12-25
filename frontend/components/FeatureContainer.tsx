import React from "react";
import { Container } from "@material-ui/core";

interface Props {
  children: React.ReactNode;
}

export default function FeatureContainer({ children }: Props): JSX.Element {
  return (
    <Container   >
      {children ?? ''}</Container>
  );
}

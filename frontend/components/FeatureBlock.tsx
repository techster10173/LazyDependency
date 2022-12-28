import React from "react";
import { Avatar, Grid, Typography } from "@mui/material";

interface Props {
  title: string;
  content: JSX.Element;
  icon?: JSX.Element;
}

export default function FeatureBlock({
  title,
  content,
  icon,
}: Props): JSX.Element {
  return (
   
    <Grid item xs={12} sm={6} md={4}  >
   
      {icon !== undefined ? (
        icon
      ) : null}
      <Typography variant="h3" >
        {title}
      </Typography>
      <Typography variant="body2">{content}</Typography>
      
    </Grid>
    
  );
}
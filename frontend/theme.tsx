import { createTheme } from "@material-ui/core/styles";


// Create a theme instance.
const theme = createTheme({
  palette: {
    primary: {
      main: "#8E44AD",
    },
    secondary: {
      main: "#9575CD",
    },
  },
  typography: {
    fontWeightMedium: 600,
    // fontFamily: "'Nunito', sans-serif",
    fontFamily: "'Open Sans', sans-serif",
    fontSize: 17,
    h1: {
      fontSize: "2.7rem",
      fontWeight: 700,
      color: "#000000",
      marginTop: "2%",
    },
    h2: {
      fontSize: "2.7rem",
      fontWeight: 700,
      color: "#000000",
    },
    h3: {
      fontSize: "1.5rem",
      fontWeight: 700,
      lineHeight: 1.25,
      color: "#000000",
      marginTop: "3%",
      marginBottom: "3%"
    },
    h4: {
      fontSize: "1.3rem",
      fontWeight: 600,
      lineHeight: 1.25,
      color: "#000000",
    },
    body1: {
      fontSize: "1.15rem",
      marginTop: "0%",
      alignItems: "left",
      marginBottom: "2%"
    },
    body2: {
      fontSize: "1.0rem",
      marginTop: "1%",
      marginBottom: "2%",
    },
    caption: {
      fontSize: "0.9rem",
      fontWceight: 400,
      // lineHeight: 1.5,
      color: "#85858B",
    },
    subtitle1: {
      fontSize: "1.0rem",
      fontWceight: 400,
      lineHeight: 1.5,
      color: "#85858B",
    },
    subtitle2: {
      fontSize: "0.8rem",
      fontWeight: 400,
      lineHeight: 1.5,
      color: "#BBBBC1",
    },
    
  },
  
  

});

export default theme;
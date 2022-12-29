
import { Box } from "@mui/material";
import Footer from "./Footer";
import Header from "./Header";


type Props = {
  children: React.ReactNode;
};

export default function Layout({ children }: Props): JSX.Element {
  return (
    <Box
      
    >
      <Header />
      <main>{children}</main>
      <Footer />
    </Box>
  );
}
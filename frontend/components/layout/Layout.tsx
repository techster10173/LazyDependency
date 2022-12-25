import { Box } from "@material-ui/core";
import Footer from "../layout/Footer";
import Header from "../layout/Header";

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
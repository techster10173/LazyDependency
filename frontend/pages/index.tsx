import { Button, Container, Box } from "@mui/material";
import CodeIcon from "@mui/icons-material/Code";
import StarIcon from "@mui/icons-material/Star";
import PeopleIcon from "@mui/icons-material/People";
import Page from "../components/layout/Page";
import HeroSection from "../components/HeroSection";
import React from "react";
import FeatureContainer from "../components/FeatureContainer";
import FeatureBlocksContainer from "../components/FeatureBlocksContainer";
import FeatureBlock from "../components/FeatureBlock";
import ContactSection from "../components/ContactSection";
import Link from "next/link";

export default function HomePage(): JSX.Element {
  return (
    <Page maxWidth={false}>
      <HeroSection
        title="Lazy Dependency"
        subtitle="Forget Dependency Management Problems."
        image="/assets/city.svg"
      >
         
        <Button href="/api/auth/login" variant="contained"   >Get started
        </Button>
        
        {/* <Link href="/ios" passHref> */}
        <Button disableElevation >
          
        </Button>
        {/* </Link> */}
      </HeroSection>
      
      <FeatureContainer >
        <FeatureBlocksContainer>
          <FeatureBlock
            title="Innovative"
            icon={<PeopleIcon />}
            content={
              <>
                Simplify the process of managing the external libraries and frameworks that your project depends on.
              </>
            }
          />
          <FeatureBlock
            title="Simple"
            icon={<StarIcon />}
            content={
              <>
                Easily manage multiple environments, such as development, staging, and production.
              </>
            }
          />
          <FeatureBlock
            title="Scalable"
            icon={<CodeIcon />}
            content={
              <>
                Handle a large number of dependencies without becoming unwieldy or inefficient.
              </>
            }
          />
          
        </FeatureBlocksContainer>
      </FeatureContainer>
      
      
      
      <Box >
        <ContactSection  />
      </Box>
    </Page>
  );
}
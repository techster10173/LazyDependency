import { Button, Container, Box } from "@material-ui/core";
import CodeIcon from "@material-ui/icons/Code";
import StarIcon from "@material-ui/icons/Star";
import PeopleIcon from "@material-ui/icons/People";
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
        <Link href="/api/auth/login" passHref> 
        <Button variant="contained"   >Get started
        </Button>
        </Link>
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
import { Button, Container, Box } from "@material-ui/core";
import EmailIcon from "@material-ui/icons/Email";
import PersonIcon from "@material-ui/icons/Person";
import LockIcon from "@material-ui/icons/Lock";
import Page from "../components/layout/Page";
import HeroSection from "../components/HeroSection";
import React from "react";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Input from "@material-ui/core/Input";
import FormHelperText from "@material-ui/core/FormHelperText";

export default function RegisterPage(): JSX.Element {
  return (
    <Page maxWidth={false}>
      <HeroSection
        title="Account Login"
        image="/assets/city.svg"
      >
        <form>
        <FormControl>
            <InputLabel htmlFor="username">Username</InputLabel>
            <Input
              id="username"
              type="text"
              aria-describedby="username-helper-text"
            />
            <FormHelperText id="username-helper-text">Enter your username.</FormHelperText>
          </FormControl>
          <FormControl>
            <InputLabel htmlFor="password">Password</InputLabel>
            <Input
              id="password"
              type="password"
              aria-describedby="password-helper-text"
            />
            <FormHelperText id="password-helper-text">Enter your password.</FormHelperText>
          </FormControl>
          
          <Button variant="contained">
            Login
          </Button>
        </form>
      </HeroSection>
    </Page>
  );
}
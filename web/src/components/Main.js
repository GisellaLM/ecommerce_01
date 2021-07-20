import React from "react";
import Container from "@material-ui/core/Container";
import { Grid } from "@material-ui/core";
import Menu from "./Menu";
import MainImage from "./MainImage";
import SectionOne from "./SectionOne";
import Box from "@material-ui/core/Box";

export default function Main() {
  return (
    <>
      <div>
        <Container maxWidth={false} disableGutters={true} style={{ margin: 0 }}>
          <header>
            <Grid container>
              <Grid item xs={12} md={12} lg={12}>
                <Menu></Menu>
              </Grid>
              <Grid item xs={12} md={12} lg={12}>
                <MainImage></MainImage>
              </Grid>
              <Grid item xs={12} md={12} lg={12}>
                <Box mt={4}>
                  <SectionOne></SectionOne>
                </Box>
              </Grid>
            </Grid>
          </header>
          <footer></footer>
        </Container>
      </div>
    </>
  );
}

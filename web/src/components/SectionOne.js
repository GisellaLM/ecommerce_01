import React from "react";
import { makeStyles, useTheme } from "@material-ui/core/styles";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Typography from "@material-ui/core/Typography";
import { Grid } from "@material-ui/core";
import Image from "../img/card1.jpg";
import Box from "@material-ui/core/Box";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
    height: "150px",
    width: "500px",
  },
  details: {
    display: "flex",
    flexDirection: "column",
    backgroundColor: "#b1bf9d",
    fontFamily: "Open Sans",
  },
  content: {
    flex: "1 0 auto",
  },
  cover: {
    width: 200,
    borderLeft: "solid 11px #b1bf9d",
    borderRadius: "88px 0px 0px 88px",
  },
  controls: {
    display: "flex",
    alignItems: "center",
    paddingLeft: theme.spacing(1),
    paddingBottom: theme.spacing(1),
  },
  playIcon: {
    height: 38,
    width: 38,
  },
  divImgcolor: {
    backgroundColor: "#b1bf9d",
    display: "flex",
  },
  text: {
    fontFamily: "Open Sans",
  },
}));

export default function SectionOne() {
  const classes = useStyles();

  return (
    <>
      <Grid container>
        <Grid item md={1}></Grid>
        <Grid item md={5}>
          <Card className={classes.root}>
            <div className={classes.details}>
              <CardContent className={classes.content}>
                <Typography component="h5" variant="h5">
                  SPRING
                </Typography>
                <Typography component="h3" variant="h3">
                  2021
                </Typography>
                <Typography component="h5" variant="h5" noWrap={false}>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </Typography>
                <Typography variant="subtitle1" color="textSecondary">
                  Mac Miller
                </Typography>
              </CardContent>
            </div>
            <div className={classes.divImgcolor}>
              <CardMedia
                className={classes.cover}
                image={Image}
                title="Live from space album cover"
              />
            </div>
          </Card>
        </Grid>

        <Grid item md={5}>
          <Card className={classes.root}>
            <div className={classes.details}>
              <CardContent className={classes.content}>
                <Typography component="h5" variant="h5">
                  SPRING
                </Typography>
                <Typography component="h5" variant="h5">
                  2021
                </Typography>
                <Typography component="h5" variant="h5" noWrap={false}>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </Typography>
                <Typography variant="subtitle1" color="textSecondary">
                  Mac Miller
                </Typography>
              </CardContent>
            </div>
            <div className={classes.divImgcolor}>
              <CardMedia
                className={classes.cover}
                image={Image}
                title="Live from space album cover"
              />
            </div>
          </Card>
        </Grid>
        <Grid item md={1}></Grid>
      </Grid>
    </>
  );
}

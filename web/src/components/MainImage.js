import React from "react";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Image from "../img/background_1.png";
import PlayCircleOutlineIcon from "@material-ui/icons/PlayCircleOutline";
import { Grid } from "@material-ui/core";
import Box from "@material-ui/core/Box";

const useStyles = makeStyles((theme) => ({
  buttonStyle: {
    margin: theme.spacing(1),
    borderRadius: "30px",
    paddingLeft: "50px",
    paddingRight: "50px",
    paddingTop: "10px",
    paddingBottom: "10px",
    fontWeight: "bold",
    textTransform: "none",
  },
  bNowBrowse: {
    color: "#FFFFFF",
    backgroundColor: "#B1BF9D",
  },
  bWatchVideo: {
    color: "#B1BF9D",
    borderColor: "#B1BF9D",
  },
  divImage: {
    backgroundImage: `url(${Image})`,
    height: "450px",
    width: "100%",
    backgroundSize: "cover",
    backgroundRepeat: "no-repeat",
  },
  text: {
    color: "#444444",
    fontFamily: "Open Sans",
  },
  tFonth3: {
    fontWeight: "600",
  },
  tFonth5: {
    fontWeight: "400",
  },
}));

export default function () {
  const classes = useStyles();

  return (
    <>
      <div className={classes.divImage}>
        <Grid container>
          <Grid item md={1}></Grid>
          <Grid item md={4} alignItems="center">
            <Box mt={6}>
              <div>
                <Typography
                  className={`${classes.text} ${classes.tFonth5}`}
                  variant="h5"
                >
                  Are you ready to
                </Typography>
                <Typography
                  className={`${classes.text} ${classes.tFontw3}`}
                  variant="h3"
                >
                  New Collection
                </Typography>
                <Box mt={3} mb={3}>
                  <Typography variant="p" className={`${classes.text}`}>
                    Lorem Ipsum is simply dummy text of the printing and
                    typesetting industry. Lorem Ipsum has been the industry's
                    standard dummy text ever since the 1500s, when an unknown
                    printer took a galley of type and scrambled it to make a
                    type specimen book.
                  </Typography>
                </Box>
              </div>
            </Box>
            <div>
              <Button
                variant="contained"
                className={`${classes.buttonStyle} ${classes.bNowBrowse}`}
              >
                Now Browse
              </Button>
              <Button
                startIcon={<PlayCircleOutlineIcon />}
                variant="outlined"
                className={`${classes.buttonStyle} ${classes.bWatchVideo}`}
              >
                Watch Video
              </Button>
            </div>
          </Grid>
        </Grid>
        <Grid container>
          <Grid item md={6}></Grid>
        </Grid>
      </div>
    </>
  );
}

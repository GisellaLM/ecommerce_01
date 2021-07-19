import React from "react";
import { Link } from "@material-ui/core";
import Button from "@material-ui/core/Button";
import Typography from "@material-ui/core/Typography";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";
import { makeStyles } from "@material-ui/core/styles";
import InputBase from "@material-ui/core/InputBase";

const useStyles = makeStyles((theme) => ({
  grow: {
    flexGrow: 4,
  },
  menuButton: {
    marginRight: theme.spacing(2),
    fontFamily: "Open Sans",
  },
  title: {
    display: "none",
    [theme.breakpoints.up("sm")]: {
      display: "block",
    },
  },
  search: {
    position: "relative",
    borderRadius: theme.shape.borderRadius,
    flexGrow: 4,
    marginRight: theme.spacing(2),
    marginLeft: 0,
    width: "100%",
    [theme.breakpoints.up("sm")]: {
      marginLeft: theme.spacing(3),
      width: "auto",
    },
  },
  searchIcon: {
    padding: theme.spacing(0, 2),
    height: "100%",
    position: "absolute",
    pointerEvents: "none",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
  },
  inputRoot: {
    color: "inherit",
  },
  inputInput: {
    padding: theme.spacing(1, 1, 1, 0),
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
    transition: theme.transitions.create("width"),
    width: "100%",
    [theme.breakpoints.up("md")]: {
      width: "20ch",
    },
  },
  sectionDesktop: {
    display: "none",
    [theme.breakpoints.up("md")]: {
      display: "flex",
    },
  },
  sectionMobile: {
    display: "flex",
    [theme.breakpoints.up("md")]: {
      display: "none",
    },
  },
  textButton: {
    fontFamily: "Open Sans",
    color: "#444444",
  },
}));

export default function Menu() {
  const classes = useStyles();

  return (
    <>
      <div className={classes.grow}>
        <AppBar color="transparent" position="static">
          <Toolbar>
            <div>
              <Button className={classes.menuButton} color="inherit">
                {" "}
                <MenuIcon></MenuIcon>
                <div> All category</div>
              </Button>
            </div>

            <div className={classes.search}>
              <div className={classes.searchIcon}>
                <SearchIcon />
              </div>
              <InputBase
                placeholder="Search…"
                classes={{
                  root: classes.inputRoot,
                  input: classes.inputInput,
                }}
                inputProps={{ "aria-label": "search" }}
              />
            </div>

            <div className={classes.grow}>
              <Button className={classes.textButton} color="default">
                Home
              </Button>
              <Button className={classes.textButton} color="inherit">
                Shop
              </Button>
              <Button className={classes.textButton} color="inherit">
                Products
              </Button>
              <Button className={classes.textButton} color="inherit">
                About Us
              </Button>
              <Button className={classes.textButton} color="inherit">
                Blog
              </Button>
              <Button className={classes.textButton} color="inherit">
                FAQ
              </Button>
              <Button className={classes.textButton} color="inherit">
                Contact Us
              </Button>
            </div>
          </Toolbar>
        </AppBar>
      </div>
    </>
  );
}

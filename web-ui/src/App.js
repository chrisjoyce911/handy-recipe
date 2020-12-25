import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";

import {
  AppBar,
  Link,
  Toolbar,
  Paper,
  IconButton,
  Typography,
  InputBase,
} from "@material-ui/core";

import { fade, makeStyles } from "@material-ui/core/styles";
import MenuIcon from "@material-ui/icons/Menu";
import SearchIcon from "@material-ui/icons/Search";

import HomePage from "./pages/HomePage"
import UserPage from "./pages/UserPage"
import RecipePage from "./pages/RecipePage"

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,

  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
    display: "none",
    [theme.breakpoints.up("sm")]: {
      display: "block",
    },
  },
  search: {
    position: "relative",
    borderRadius: theme.shape.borderRadius,
    backgroundColor: fade(theme.palette.common.white, 0.15),
    "&:hover": {
      backgroundColor: fade(theme.palette.common.white, 0.25),
    },
    marginLeft: 0,
    width: "100%",
    [theme.breakpoints.up("sm")]: {
      marginLeft: theme.spacing(1),
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
    [theme.breakpoints.up("sm")]: {
      width: "12ch",
      "&:focus": {
        width: "20ch",
      },
    },
  },
}));


function Main() {
const classes = useStyles();
 return (
   <BrowserRouter>
     <div className={classes.root}>
       <AppBar position="static">
         <Toolbar>
           <IconButton
             edge="start"
             className={classes.menuButton}
             color="inherit"
             aria-label="open drawer"
           >
             <MenuIcon />
           </IconButton>
           <Typography className={classes.title} variant="h6" noWrap>
             <Link href="/" color="inherit">
               Handy Recipe
             </Link>
           </Typography>

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
         </Toolbar>
       </AppBar>
     </div>

     <Paper variant="outlined">
       <Switch>
         <Route exact path="/" component={HomePage} />
         <Route path="/recipe/:id" component={RecipePage} />
         <Route path="/:id" component={UserPage} />
       </Switch>
     </Paper>
   </BrowserRouter>
 );
};

export default Main;

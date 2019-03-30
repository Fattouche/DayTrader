import React, {Component} from 'react';
import PropTypes from 'prop-types';
import { MuiThemeProvider, createMuiTheme, withStyles } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';
import Hidden from '@material-ui/core/Hidden';
import Navigator from './Navigator';
import MyProfile from './MyProfile';
import Header from './Header';
import Sell from './Sell';
import Buy from './Buy';
import Browse from './Browse';
import Triggers from './Triggers';
import AddFunds from './AddFunds';
import Dumplog from './Dumplog';


let theme = createMuiTheme({
  typography: {
    useNextVariants: true,
    h5: {
      fontWeight: 500,
      fontSize: 26,
      letterSpacing: 0.5,
    },
  },
  palette: {
    primary: {
      light: '#63ccff',
      main: '#009be5',
      dark: '#006db3',
    },
  },
  shape: {
    borderRadius: 8,
  },
});

theme = {
  ...theme,
  overrides: {
    MuiDrawer: {
      paper: {
        backgroundColor: '#18202c',
      },
    },
    MuiButton: {
      label: {
        textTransform: 'initial',
      },
      contained: {
        boxShadow: 'none',
        '&:active': {
          boxShadow: 'none',
        },
      },
    },
    MuiTabs: {
      root: {
        marginLeft: theme.spacing.unit,
      },
      indicator: {
        height: 3,
        borderTopLeftRadius: 3,
        borderTopRightRadius: 3,
        backgroundColor: theme.palette.common.white,
      },
    },
    MuiTab: {
      root: {
        textTransform: 'initial',
        margin: '0 16px',
        minWidth: 0,
        [theme.breakpoints.up('md')]: {
          minWidth: 0,
        },
      },
      labelContainer: {
        padding: 0,
        [theme.breakpoints.up('md')]: {
          padding: 0,
        },
      },
    },
    MuiIconButton: {
      root: {
        padding: theme.spacing.unit,
      },
    },
    MuiTooltip: {
      tooltip: {
        borderRadius: 4,
      },
    },
    MuiDivider: {
      root: {
        backgroundColor: '#404854',
      },
    },
    MuiListItemText: {
      primary: {
        fontWeight: theme.typography.fontWeightMedium,
      },
    },
    MuiListItemIcon: {
      root: {
        color: 'inherit',
        marginRight: 0,
        '& svg': {
          fontSize: 20,
        },
      },
    },
    MuiAvatar: {
      root: {
        width: 32,
        height: 32,
      },
    },
  },
  props: {
    MuiTab: {
      disableRipple: true,
    },
  },
  mixins: {
    ...theme.mixins,
    toolbar: {
      minHeight: 48,
    },
  },
};

const drawerWidth = 256;

const styles = {
  root: {
    display: 'flex',
    minHeight: '100vh',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
  appContent: {
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
  },
  mainContent: {
    flex: 1,
    padding: '48px 36px 0',
    background: '#eaeff1',
  },
};

class Landing extends Component {
    constructor(props){
        super(props)
        this.contentMap = {
          'My Profile': <MyProfile userInfo={props.userInfo}/>, 
          'Buy': <Buy userId={props.userInfo.getUserId()}/>, 
          'Sell': <Sell userId={props.userInfo.getUserId()}/>, 
          'Browse': <Browse userId={props.userInfo.getUserId()}/>,
          'Triggers': <Triggers userId={props.userInfo.getUserId()}/>,
          'Add Funds': <AddFunds userId={props.userInfo.getUserId()}/>,
          'Dumplog': <Dumplog userId={props.userInfo.getUserId()}/>
        }

        this.state = {
            mobileOpen: false,
            content: 'My Profile',
            userInfo: props.userInfo
        };

        this.handler = props.handler
        this.classes = props.classes
        this.showContentAndHeader = this.showContentAndHeader.bind(this)
        this.logout = this.logout.bind(this)
    }
 
  handleDrawerToggle = () => {
    this.setState({ mobileOpen: !this.state.mobileOpen });
  };

  showContentAndHeader(contentId){
    this.setState({
        content: contentId
    })
  }

  logout(){
    this.handler()
  }

  render() {
    return (
      <MuiThemeProvider theme={theme}>
        <div className={this.classes.root}>
          <CssBaseline />
          <nav className={this.classes.drawer}>
            <Hidden smUp implementation="js">
              <Navigator
                PaperProps={{ style: { width: drawerWidth } }}
                variant="temporary"
                open={this.state.mobileOpen}
                onClose={this.handleDrawerToggle}
              />
            </Hidden>
            <Hidden xsDown implementation="css">
              <Navigator PaperProps={{ style: { width: drawerWidth } }} handler={[this.showContentAndHeader,this.logout]} contentId={this.state.content}/>
            </Hidden>
          </nav>
          <div className={this.classes.appContent}>
            <Header onDrawerToggle={this.handleDrawerToggle} title={this.state.content}/>
            <main className={this.classes.mainContent}>
              {this.contentMap[this.state.content]}
            </main>
          </div>
        </div>
      </MuiThemeProvider>
    );
  }
}

Landing.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Landing);
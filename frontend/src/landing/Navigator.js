import React, {Component} from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import { withStyles } from '@material-ui/core/styles';
import Divider from '@material-ui/core/Divider';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import HomeIcon from '@material-ui/icons/Home';
import PeopleIcon from '@material-ui/icons/People';
import PublicIcon from '@material-ui/icons/Public';
import SettingsPower from '@material-ui/icons/SettingsPower';
import AttachMoney from '@material-ui/icons/AttachMoney';
import MoneyOff from '@material-ui/icons/MoneyOff';
import SettingsRemote from '@material-ui/icons/SettingsRemote';
import AccessibleForward from '@material-ui/icons/AccessibleForward';


const categories = [
  {
    id: 'Manage',
    children: [
      { id: 'My Profile', icon: <PeopleIcon /> },
      { id: 'Add Funds', icon: <AccessibleForward/> },
      { id: 'Buy', icon: <AttachMoney /> },
      { id: 'Sell', icon: <MoneyOff /> },
      { id: 'Browse', icon: <PublicIcon /> },
      { id: 'Triggers', icon: <SettingsRemote /> },
    ],
  },
  {
    id: 'Account',
    children: [
      { id: 'Logout', icon: <SettingsPower /> },
    ],
  },
];

const styles = theme => ({
  categoryHeader: {
    paddingTop: 16,
    paddingBottom: 16,
  },
  categoryHeaderPrimary: {
    color: theme.palette.common.white,
  },
  item: {
    paddingTop: 4,
    paddingBottom: 4,
    color: 'rgba(255, 255, 255, 0.7)',
  },
  itemCategory: {
    backgroundColor: '#232f3e',
    boxShadow: '0 -1px 0 #404854 inset',
    paddingTop: 16,
    paddingBottom: 16,
  },
  firebase: {
    fontSize: 24,
    fontFamily: theme.typography.fontFamily,
    color: theme.palette.common.white,
  },
  itemActionable: {
    '&:hover': {
      backgroundColor: 'rgba(255, 255, 255, 0.08)',
    },
  },
  itemActiveItem: {
    color: '#4fc3f7',
  },
  itemPrimary: {
    color: 'inherit',
    fontSize: theme.typography.fontSize,
    '&$textDense': {
      fontSize: theme.typography.fontSize,
    },
  },
  textDense: {},
  divider: {
    marginTop: theme.spacing.unit * 2,
  },
});

class Navigator extends Component {
    constructor(props){
        super(props)
        this.classes = props.classes;
        this.props = props
        this.changeContent = props.handler[0];
        this.logout = props.handler[1];
    }

render(){
  return (
    <Drawer variant="permanent" {...this.props}>
      <List disablePadding>
        <ListItem className={classNames(this.classes.firebase, this.classes.item, this.classes.itemCategory)}>
          Daytrader
        </ListItem>
        <ListItem className={classNames(this.classes.item, this.classes.itemCategory)}>
          <ListItemIcon>
            <HomeIcon />
          </ListItemIcon>
          <ListItemText
            classes={{
              primary: this.classes.itemPrimary,
            }}
          >
            Menu Overview
          </ListItemText>
        </ListItem>
        {categories.map(({ id, children }) => (
          <React.Fragment key={id}>
            <ListItem className={this.classes.categoryHeader}>
              <ListItemText
                classes={{
                  primary: this.classes.categoryHeaderPrimary,
                }}
              >
                {id}
              </ListItemText>
            </ListItem>
            {children.map(({ id: childId, icon }) => (
              <ListItem
                button
                dense
                key={childId}
                className={classNames(
                  this.classes.item,
                  this.classes.itemActionable,
                  this.props.contentId === childId && this.classes.itemActiveItem,
                )}
                onClick={() => {childId === "Logout" ? this.logout() : this.changeContent(childId)}}
              >
                <ListItemIcon>{icon}</ListItemIcon>
                <ListItemText
                  classes={{
                    primary: this.classes.itemPrimary,
                    textDense: this.classes.textDense,
                  }}
                >
                  {childId}
                </ListItemText>
              </ListItem>
            ))}
            <Divider className={this.classes.divider} />
          </React.Fragment>
        ))}
      </List>
    </Drawer>
  );
}

}

Navigator.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Navigator);

import { Navbar, Group, Code, ScrollArea, createStyles, rem, NavLink, LoadingOverlay } from '@mantine/core';
import { IconActivity, IconDashboard, IconUserCircle, IconUsers } from '@tabler/icons-react';
import { Logo } from './Logo';
import NavHeader from './NavHeader';
import { NavFooter } from './NavFooter';
import { NavLink as RouterNavLink } from "react-router-dom"
import React from 'react';
import { useAuthUserQuery } from '@/redux/services/auth';


// const navItems = [
//   { label: 'Dashboard', icon: IconLock, link: '/dashboard' },
//   { label: 'Users', icon: IconCalendarStats, link: '/users' },
//   {
//     label: 'Market news',
//     icon: IconNotes,
//     initiallyOpened: false,
//     links: [
//       { label: 'Overview', link: '/Overview' },
//       { label: 'Forecasts', link: '/Forecasts' },
//       { label: 'Outlook', link: '/Outlook' },
//       { label: 'Real time', link: '/Real time' },
//     ],
//   },

// ];

const useStyles = createStyles((theme) => ({
  navbar: {
    backgroundColor: theme.colors.orange[8],
    paddingBottom: 0,
    height: '100vh'
  },

  header: {
    padding: theme.spacing.md,
    paddingTop: 0,
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
    color: theme.white,
    borderBottom: `${rem(1)} solid ${theme.colors.gray[3]
      }`,
  },

  links: {
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
    // color: theme.white,
  },

  linksInner: {
    paddingTop: theme.spacing.xl,
    paddingBottom: theme.spacing.xl,
    // color: theme.white,
  },

  footer: {
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
    borderTop: `${rem(1)} solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
  },
}));




export function NavbarNested() {
  const { classes } = useStyles();
  const { data: userAuth, isLoading } = useAuthUserQuery();
  // const links = navItems.map((item) => <LinksGroup {...item} key={item.label} />);
  const permissions = userAuth?.edges?.role?.edges?.permissions?.map(permission => permission.value) || [];
  return (
    <Navbar height={800} width={{ sm: 300 }} p="md" className={classes.navbar}>
      <Navbar.Section className={classes.header}>
        <NavHeader />
      </Navbar.Section>

      <Navbar.Section grow className={classes.links} component={ScrollArea}>
        {isLoading ?
          <LoadingOverlay loaderProps={{ size: 'sm', color: 'orange', variant: 'bars' }} visible={isLoading} />
          :
          <>
            {/* <CustomNavLink label="dashboard" icon={<IconActivity />} > */}
              <RouterNavLink to="/dashboard" className={" no-underline"} >
                {({ isActive }) => (
                  <CustomNavLink label="dashboard" icon={<IconDashboard/>} isActive={isActive} />
                )}
              </RouterNavLink>
            {/* </CustomNavLink> */}
            <RouterNavLink to="/users" className={permisionsVerify(["create_user"], permissions) ? "no-underline" : "hidden"}>
              {({ isActive }) => (
                <CustomNavLink label="users" icon={<IconUsers />} isActive={isActive} />
              )}
            </RouterNavLink>
          </>
        }

      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <NavFooter/>
      </Navbar.Section>
    </Navbar>
  );
}

export const Head = () => {
  const { classes } = useStyles();
  return (<Navbar.Section className={classes.header}>
    <Group position="apart">
      <Logo color='white' width={rem(120)} />
      <Code sx={{ fontWeight: 700 }}>v.1</Code>
    </Group>
  </Navbar.Section>
  )
}

interface CustomNavLinkInterface {
  isActive?: boolean;
  icon: React.ReactElement;
  label: string;
  children?: React.ReactElement | undefined;
}
const CustomNavLink: React.FC<CustomNavLinkInterface> = ({ isActive = false, icon, label, children }) => {
  return <>
    <NavLink
      label={label}
      icon={icon}
      // active={isActive}
      className={`${isActive ? 'bg-orange-700' : 'bg-orange-600'} text-orange-50 hover:bg-orange-700 duration-300`}
    >
      {children}
    </NavLink>
  </>
}

const permisionsVerify = (requiredPermission: string[], authPermission: string[]): boolean => {
  for (const element of requiredPermission) {
    if (authPermission.includes(element)) return true
  }
  return false
}
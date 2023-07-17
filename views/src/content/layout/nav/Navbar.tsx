import { Navbar, Group, Code, ScrollArea, createStyles, rem, MantineProvider } from '@mantine/core';
import {
  IconNotes,
  IconCalendarStats,
  IconGauge,
  IconPresentationAnalytics,
  IconFileAnalytics,
  IconAdjustments,
  IconLock,
} from '@tabler/icons-react';
import { LinksGroup } from './NavbarLinksGroups';
import { Logo } from './Logo';
import NavHeader from './NavHeader';
import { NavFooter } from './NavFooter';
import { useState } from 'react';


const navItems = [
  { label: 'Dashboard', icon: IconGauge, link:'/dashboard' },
  {
    label: 'Market news',
    icon: IconNotes,
    initiallyOpened: false,
    links: [
      { label: 'Overview', link: '/' },
      { label: 'Forecasts', link: '/' },
      { label: 'Outlook', link: '/' },
      { label: 'Real time', link: '/' },
    ],
  },
  {
    label: 'Market news',
    icon: IconNotes,
    initiallyOpened: false,
    links: [
      { label: 'Overview', link: '/' },
      { label: 'Forecasts', link: '/' },
      { label: 'Outlook', link: '/' },
      { label: 'Real time', link: '/' },
    ],
  }
];

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
  const links = navItems.map((item) => <LinksGroup {...item} key={item.label} />);

  return (
    <Navbar height={800} width={{ sm: 300 }} p="md" className={classes.navbar}>
      <Navbar.Section className={classes.header}>
        <NavHeader/>
      </Navbar.Section>

      <Navbar.Section grow className={classes.links} component={ScrollArea}>
        <div className={classes.linksInner}>{links}</div>
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <NavFooter
          image="https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80"
        />
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
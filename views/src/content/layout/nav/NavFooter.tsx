import { useLoginUserMutation } from '@/redux/services/auth';
import {
  UnstyledButton,
  UnstyledButtonProps,
  Group,
  Avatar,
  Text,
  createStyles,
} from '@mantine/core';
import { IconChevronRight } from '@tabler/icons-react';

const useStyles = createStyles((theme) => ({
  user: {
    display: 'block',
    width: '100%',
    padding: theme.spacing.md,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.colors.orange[0],

    '&:hover': {
      backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.orange[0],
      color: theme.colorScheme === 'dark' ? theme.colors.dark[0] : theme.colors.orange[8],

    },
  },
  text: {
    opacity: 0.8,
  }
}));

interface UserButtonProps extends UnstyledButtonProps {
  image: string;
}

export function NavFooter({ image, ...others }: UserButtonProps) {
  
  const { classes } = useStyles();
  const {data:user}=useLoginUserMutation();

  return (
    <UnstyledButton className={classes.user} {...others}>
      <Group>
        <Avatar src={image} radius="xl" />

        <div style={{ flex: 1 }}>
          <Text size="sm" weight={500}>
            {user?.name}
          </Text>

          <Text size="xs" className={classes.text}>
            {user?.email}
          </Text>
        </div>

        {user?.icon || <IconChevronRight size="0.9rem" stroke={1.5} />}
      </Group>
    </UnstyledButton>
  );
}
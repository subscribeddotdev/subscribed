import { Box, Heading, Text } from "@radix-ui/themes";

interface Props {
  title: string;
  label?: string;
}

export function PageTitle({ title, label }: Props) {
  return (
    <Box>
      {label && (
        <Text color="gray" weight="bold" size="2">
          {label}
        </Text>
      )}
      <Heading>{title}</Heading>
    </Box>
  );
}

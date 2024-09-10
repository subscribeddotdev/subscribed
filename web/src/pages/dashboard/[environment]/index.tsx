import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { Flex } from "@radix-ui/themes";
import { GetServerSideProps } from "next";

interface Props {
  token: string;
}

export default function DashboardHomePage({ token }: Props) {
  return (
    <LayoutDashboard>
      <PageMeta title="Dashboard" />
      <Flex justify="between" mb="4">
        <PageTitle title="Get started" />
      </Flex>
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  return {
    props: { token: "" },
  };
};

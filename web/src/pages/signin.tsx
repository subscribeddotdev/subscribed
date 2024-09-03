import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutAuth } from "@@/common/layouts/LayoutAuth/LayoutAuth";
import { SignInForm } from "@@/modules/Auth/SignInForm/SignInForm";
import { GetServerSideProps } from "next";

interface Props {}

export default function SignInPage({}: Props) {
  return (
    <LayoutAuth>
      <PageMeta title="Sign in" />
      <SignInForm />
    </LayoutAuth>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  return {
    props: {},
  };
};

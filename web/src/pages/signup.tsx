import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutAuth } from "@@/common/layouts/LayoutAuth/LayoutAuth";
import { SignUpForm } from "@@/modules/Auth/SignUpForm/SignUpForm";

interface Props {}

export default function SignUpPage({}: Props) {
  return (
    <LayoutAuth>
      <PageMeta title="Sign up" />
      <SignUpForm />
    </LayoutAuth>
  );
}

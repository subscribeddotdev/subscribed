import { Alert } from "@@/common/components/Alert/Alert";
import { Button } from "@@/common/components/Button/Button";
import { Input } from "@@/common/components/Input/Input";
import { createApiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { paths } from "@@/constants";
import { Box, Flex, Grid, Heading, Link, Text } from "@radix-ui/themes";
import { useFormik } from "formik";
import { useRouter } from "next/router";
import { useState } from "react";
import * as yup from "yup";
import styles from "./SignUpForm.module.css";

interface Props {}

export function SignUpForm({}: Props) {
  const [error, setError] = useState("");
  const router = useRouter();

  const f = useFormik({
    validationSchema,
    validateOnChange: false,
    initialValues: { first_name: "", last_name: "", email: "", password: "" },
    onSubmit: async (values, { setSubmitting }) => {
      try {
        setError("");
        await createApiClients("").Auth.signUp(values);
        await router.push(paths.signin + "?signup-succeeded=1");
      } catch (error) {
        setError(getApiError(error));
      }
    },
  });

  return (
    <form className={styles.root} onSubmit={f.handleSubmit}>
      <Heading mb="4">Sign up</Heading>

      <Flex direction="column" gap="2" mb="4">
        <Grid gap="2" columns="2">
          <Input name="first_name" required onChange={f.handleChange} label="First name" error={f.errors.first_name} />
          <Input name="last_name" required onChange={f.handleChange} label="Last name" error={f.errors.last_name} />
        </Grid>
        <Input name="email" required type="email" onChange={f.handleChange} label="E-mail" error={f.errors.email} />
        <Input name="password" label="Password" type="password" onChange={f.handleChange} error={f.errors.password} />
      </Flex>

      <Button mb="4" type="submit" disabled={f.isSubmitting} loading={f.isSubmitting}>
        Create account
      </Button>

      <Box mb="4">
        <Text size="2">
          Already have an account? <Link href={paths.signin}>Sign in.</Link>
        </Text>
      </Box>

      {error && <Alert color="red">{error}</Alert>}
    </form>
  );
}

function validationSchema() {
  return yup.object().shape({
    email: yup.string().email().required(),
    password: yup.string().required(),
  });
}

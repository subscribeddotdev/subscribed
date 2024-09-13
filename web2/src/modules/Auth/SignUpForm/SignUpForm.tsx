import { Alert } from "@@/common/components/Alert/Alert";
import { Button } from "@@/common/components/Button/Button";
import { Input } from "@@/common/components/Input/Input";
import {
  createApiClients,
  getApiError,
} from "@@/common/libs/backendapi/browser";
import { paths } from "@@/constants";
import { Box, Flex, Grid, Heading, Link, Text } from "@radix-ui/themes";
import { useFormik } from "formik";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import * as yup from "yup";
import styles from "./SignUpForm.module.css";

export function SignUpForm() {
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const f = useFormik({
    validationSchema,
    validateOnChange: false,
    initialValues: { first_name: "", last_name: "", email: "", password: "" },
    onSubmit: async (values) => {
      try {
        setError("");
        await createApiClients("").Auth.signUp(values);
        navigate(paths.signin + "?signup-succeeded=1");
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
          <Input
            data-testid="SignUpForm_Inp_FirstName"
            name="first_name"
            required
            onChange={f.handleChange}
            label="First name"
            error={f.errors.first_name}
          />
          <Input
            data-testid="SignUpForm_Inp_LastName"
            name="last_name"
            required
            onChange={f.handleChange}
            label="Last name"
            error={f.errors.last_name}
          />
        </Grid>
        <Input
          data-testid="SignUpForm_Inp_Email"
          name="email"
          required
          type="email"
          onChange={f.handleChange}
          label="E-mail"
          error={f.errors.email}
        />
        <Input
          data-testid="SignUpForm_Inp_Password"
          name="password"
          label="Password"
          type="password"
          onChange={f.handleChange}
          error={f.errors.password}
        />
      </Flex>

      <Button
        data-testid="SignUpForm_Btn_CreateAccount"
        mb="4"
        type="submit"
        disabled={f.isSubmitting}
        loading={f.isSubmitting}
      >
        Create account
      </Button>

      <Box mb="4">
        <Text size="2">
          Already have an account? <Link href={paths.signin}>Sign in.</Link>
        </Text>
      </Box>

      {error && (
        <Alert data-testid="SignUpForm_Alert_Error" color="red">
          {error}
        </Alert>
      )}
    </form>
  );
}

function validationSchema() {
  return yup.object().shape({
    email: yup.string().email().required(),
    password: yup.string().required(),
  });
}

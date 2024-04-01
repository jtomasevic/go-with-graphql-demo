
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../src/schema.graphql",
  
  generates: {
    "./api/gql/index.ts": {
      documents: "./api/queries/**/*.ts",
      // preset: "client",
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-react-apollo',
        'typescript-msw'
      ],
      config: { withHooks: true }
    }
  }
};

export default config;
// import type { CodegenConfig } from '@graphql-codegen/cli';

// const config: CodegenConfig = {
//   overwrite: true,
//   schema: "../src/schema.graphql",
//   documents: ["app/queries/"],
//   generates: {
//     'app/gql/': {
//       preset: "client",
//       presetConfig: {
//         gqlTagName: "gql",
//       },
//       plugins: [
//         'typescript',
//         'typescript-operations',
//         'typescript-react-apollo',
//         'typescript-msw',
//       ],
//       config: { withHooks: true },
//     }
//   }
// };

// export default config;
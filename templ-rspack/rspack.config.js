import { defineConfig } from '@rspack/cli';
import path from 'path';
import { fileURLToPath } from 'url';
import tailwindcss from '@tailwindcss/postcss';

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  experiments: {
    css: true,
  },
  entry: {
    main: './assets/ts/main.ts',
    index: './assets/css/index.css',
  },
  module: {
    rules: [
      {
        test: /\.css$/,
        use: ["postcss-loader"],
        type: "css",
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, 'dist'),
  },
});
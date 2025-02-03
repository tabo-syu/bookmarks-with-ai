import { generate } from 'openapi-typescript-codegen';
import { dirname, join } from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const main = async () => {
  await generate({
    input: join(__dirname, '../../openapi.yaml'), // Updated path to root directory
    output: join(__dirname, '../src/api'),
    client: 'react-query',
    httpClient: 'axios',
  });
};

main().catch(console.error);
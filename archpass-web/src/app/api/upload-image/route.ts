import { Uploader } from "@irys/upload";
import { BaseEth } from "@irys/upload-ethereum";
import type { NextRequest } from 'next/server';

const getIrysUploader = async () => {
  const rpcURL = process.env.NEXT_PUBLIC_RPC_URL;
  const irysUploader = await Uploader(BaseEth)
    .withWallet(process.env.PRIVATE_KEY)
    .withRpc(rpcURL)
    .devnet();

  return irysUploader;
};

export async function POST(request: NextRequest) {
  try {
    const authToken = request.headers.get('Authorization');
    const expectedToken = process.env.API_TOKEN;

    if (!authToken || authToken !== `Bearer ${expectedToken}`) {
      return new Response(
        JSON.stringify({
          success: false,
          error: 'Unauthorized'
        }),
        {
          status: 401,
          headers: { 'Content-Type': 'application/json' }
        }
      );
    }

    const formData = await request.formData();
    const file = formData.get('file') as File;
    const contentType = formData.get('contentType') as string;

    if (!file) {
      return new Response(
        JSON.stringify({
          success: false,
          error: 'No file provided'
        }),
        {
          status: 400,
          headers: { 'Content-Type': 'application/json' }
        }
      );
    }

    const irysUploader = await getIrysUploader();
    const tags = [{
      name: "Content-Type",
      value: contentType || file.type,
    }];

    const fileBuffer = Buffer.from(await file.arrayBuffer());
    const receipt = await irysUploader.upload(fileBuffer, { tags });

    return new Response(
      JSON.stringify({
        success: true,
        id: receipt.id,
        url: `https://gateway.irys.xyz/${receipt.id}`
      }),
      {
        headers: {
          'Content-Type': 'application/json',
        },
      }
    );

  } catch (e) {
    console.error("Error: ", e);
    return new Response(
      JSON.stringify({
        success: false,
        error: e.message
      }),
      {
        status: 500,
        headers: {
          'Content-Type': 'application/json',
        },
      }
    );
  }
}
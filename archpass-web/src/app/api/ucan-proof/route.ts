import { parse as didParse } from '@ipld/dag-ucan/did';
import type { ServiceAbility } from '@web3-storage/capabilities/types';
import * as Delegation from '@ucanto/core/delegation';
import { CarReader } from '@ipld/car';
import { create } from '@web3-storage/w3up-client';
import { Signer } from '@web3-storage/w3up-client/principal/ed25519';
import { StoreMemory } from '@web3-storage/w3up-client/stores/memory';
import type { NextRequest } from 'next/server';

async function parseProof(data) {
  const blocks = [];
  const reader = await CarReader.fromBytes(Buffer.from(data, 'base64'));
  for await (const block of reader.blocks()) {
    blocks.push(block);
  }
  return Delegation.importDAG(blocks);
}

// TODO(jhudiel): Implement authentication
export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams;
  const did = searchParams.get('did');
  if (!did) {
    return new Response(
      JSON.stringify({ message: 'Missing `did` parameter' }),
      {
        status: 400,
        headers: {
          'Content-Type': 'application/json',
        },
      },
    );
  }

  const principal = Signer.parse(process.env.W3_KEY ?? '');
  const store = new StoreMemory();
  const client = await create({ principal, store });
  const proof = await parseProof(process.env.W3_ADMIN_PROOF ?? '');
  await client.addSpace(proof);

  const audience = didParse(did);
  // Upload capabilities
  const abilities: ServiceAbility[] = [
    'space/blob/add',
    'space/index/add',
    'filecoin/offer',
    'upload/add',
  ];
  const expiration = Math.floor(Date.now() / 1000) + 60 * 60 * 24;
  const delegation = await client.createDelegation(audience, abilities, {
    expiration,
  });
  const archive = await delegation.archive();
  return new Response(archive.ok, {
    headers: {
      'Content-Type': 'application/octet-stream',
    },
  });
}

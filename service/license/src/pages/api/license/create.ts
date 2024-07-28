import { authSession } from '@/services/backend/auth';
import { findClusterByUIDAndClusterID } from '@/services/backend/db/cluster';
import {
  createLicenseRecord,
  generateLicenseToken,
  hasIssuedLicense
} from '@/services/backend/db/license';
import { getPaymentByID } from '@/services/backend/db/payment';
import { jsonRes } from '@/services/backend/response';
import { CreateLicenseParams, LicenseRecordPayload, PaymentStatus } from '@/types';
import type { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  try {
    const { orderID, clusterId } = req.body as CreateLicenseParams;
    const userInfo = await authSession(req.headers);
    if (!userInfo) {
      return jsonRes(res, { code: 401, message: 'token verify error' });
    }
    if (!orderID || !clusterId) {
      return jsonRes(res, { code: 400, message: 'Request parameter error' });
    }

    const payment = await getPaymentByID({ uid: userInfo.uid, orderID: orderID });
    if (!payment) {
      return jsonRes(res, { code: 400, message: 'No order found' });
    }
    const issuedLicense = await hasIssuedLicense({ uid: userInfo.uid, orderID: orderID });
    if (issuedLicense) {
      return jsonRes(res, { code: 400, message: 'orderID cannot be reused' });
    }
    const cluster = await findClusterByUIDAndClusterID({
      uid: userInfo.uid,
      clusterId: clusterId
    });
    if (!cluster?.kubeSystemID) {
      return jsonRes(res, {
        code: 400,
        message: 'The cluster is not activated and cannot be purchased'
      });
    }

    const _token = generateLicenseToken({
      type: 'Account',
      clusterID: cluster.kubeSystemID,
      data: { amount: payment.amount }
    });

    const record: LicenseRecordPayload = {
      uid: userInfo.uid,
      amount: payment.amount,
      token: _token,
      orderID: orderID,
      quota: payment.amount,
      payMethod: payment.payMethod,
      type: 'Account',
      clusterId: clusterId
    };

    if (payment.status !== PaymentStatus.PaymentSuccess) {
      return jsonRes(res, {
        code: 400,
        message: 'Unpaid'
      });
    }

    const result = await createLicenseRecord(record);
    return jsonRes(res, {
      data: result
    });
  } catch (error) {
    jsonRes(res, { code: 500, data: error });
  }
}

import { ClusterResult } from '@/types';
import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';

type clusterDetailState = {
  clusterDetail: ClusterResult | undefined;
  setClusterDetail: (data: ClusterResult) => void;
  clearClusterDetail: () => void;
};

export const useClusterDetail = create(
  immer<clusterDetailState>((set, get) => ({
    clusterDetail: undefined,
    setClusterDetail: (data) => {
      set({
        clusterDetail: data
      });
    },
    clearClusterDetail: () => set({ clusterDetail: undefined })
  }))
);

export default useClusterDetail;

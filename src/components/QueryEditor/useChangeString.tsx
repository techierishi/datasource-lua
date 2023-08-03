import { useCallback } from 'react';
import type { BasicQuery } from '../../types';
import type { ChangeOptions, EditorProps } from './types';
import {LuaFactory} from 'wasmoon';

type OnChangeType = (value: string) => void;

export function useChangeString(props: EditorProps, options: ChangeOptions<BasicQuery>): OnChangeType {
  const { onChange, onRunQuery, query } = props;
  const { propertyName, runQuery } = options;

  return useCallback(
    (value: string) => {
      if (!value) {
        return;
      }

      if (runQuery) {
        onRunQuery();
        (async () => {
          const factory = new LuaFactory();
          const lua = await factory.createEngine();
      
          try {
              lua.global.set('fetch', (url: string) => fetch(url));
      
              const data = await lua.doString(query.rawQuery);
              console.log("Data.... ", data)

              onChange({
                ...query,
                [propertyName]: value,
              });
          } finally {
              lua.global.close();
          }
      })();
      }
    },
    [onChange, onRunQuery, query, propertyName, runQuery]
  );
}

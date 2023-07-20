import { Paper } from '@mantine/core';
import { DataTable } from 'mantine-datatable';
import React, { useEffect, useState } from 'react';



interface HeadColumns {
  accessor: string;
  width?: number;
  title?: string;
  // render?: Function;
}
interface PaginateTableWrapper {
  headColumns: HeadColumns[] | [];
  datas: object[] | [undefined];
}

const PAGE_SIZES = [10, 15, 20];

const PaginationExampleWithPageSizeSelector: React.FC<PaginateTableWrapper> = ({ headColumns = [], datas = [] }) => {
  const [pageSize, setPageSize] = useState(PAGE_SIZES[1]);

  useEffect(() => {
    setPage(1);
  }, [pageSize]);

  const [page, setPage] = useState(1);
  const [records, setRecords] = useState(datas.slice(0, pageSize));

  useEffect(() => {
    const from = (page - 1) * pageSize;
    const to = from + pageSize;
    setRecords(datas.slice(from, to));
  }, [datas, page, pageSize]);

  return (
    <Paper shadow={"md"} sx={{ width: '100%', height: "100%", overflow: "auto", backgroundColor: 'white', boxShadow: ' 1px solid orange', padding: 3 }} >
      <DataTable
        highlightOnHover
        withBorder
        withColumnBorders
        records={records}
        columns={headColumns}
        // columns={[
        //   { accessor: 'firstName', width: 100 },
        //   { accessor: 'lastName', width: 100 },
        //   { accessor: 'email', width: '100%' },
        //   {
        //     accessor: 'birthDate',
        //     textAlignment: 'right',
        //     width: 120,
        //     // render: ({ birthDate }) => dayjs(birthDate).format('MMM D YYYY'),
        //   },
        // ]}
        totalRecords={datas.length}
        paginationColor="orange"
        recordsPerPage={pageSize}
        page={page}
        onPageChange={(p: number) => setPage(p)}
        recordsPerPageOptions={PAGE_SIZES}
        onRecordsPerPageChange={setPageSize}
      />
    </Paper>
  );
}

export default PaginationExampleWithPageSizeSelector;
<?php


namespace App\Traits;

use Illuminate\Contracts\Pagination\LengthAwarePaginator;

trait PaginateTrait
{
    /**
     * Format response paginate data.
     *
     * @param  LengthAwarePaginator $paginator
     * @return array
     */
    public function handleResponseData(LengthAwarePaginator $paginator)
    {
        return [
            'items'     => $paginator->items(),
            'total'     => $paginator->total(),
            'page'      => $paginator->currentPage(),
            'last_page' => $paginator->lastPage(),
            'count'     => $paginator->count(),
        ];
    }
}

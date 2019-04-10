-- Deploy get_latest_patchset

BEGIN;

CREATE OR REPLACE FUNCTION get_latest_patchset(p_change_id changes.id%TYPE)
RETURNS SETOF patchsets -- really just one thing, but it's useful to treat it as a table
LANGUAGE SQL STABLE
ROWS 1
AS $$
 SELECT p.id,
        p.change_id,
        p.sequence_number,
        p.submitted_at,
        p.sha,
        p.submitter_id,
        p.verified_against_sha,
        p.is_verified,
        p.status
   FROM patchsets as p
   JOIN changes AS c
     ON c.id = p.change_id
  WHERE c.id = $1
    AND p.sequence_number = c.latest_patchset;
$$;

COMMENT ON FUNCTION get_latest_patchset(p_change_id changes.id%TYPE) IS
$$Convenience function to retrieve the latest (i.e., most recently
submitted) patchset record for a given change.  This returns the
single row corresponding to that patchset directly from the
`patchsets` table; it does no "mapping" of submitter ID to a name, for
instance.

This gives us the path to easily finding out the status of a change
overall; it's just the status of the latest patchset!

This should only be used for retrieving the latest patchset of *a
single change*; there are more efficient ways to do this for a number
of changes at once.
$$;

COMMIT;

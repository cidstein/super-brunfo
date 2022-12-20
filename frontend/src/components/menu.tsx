import { makeStyles, MenuItem, Select } from "@material-ui/core";

export default function MenuBar() {
    const useStyles = makeStyles({
        form: {
          primary: {
            main: "#337AFF",
            contrastText: "#242526",
          },
          background:  "#242526",
          
        },
      });  

    const classes = useStyles();

    return (
        <form className={classes.form}>
            <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                value={''}
                label="Age"
                title="teste"
                

                
            >
                <MenuItem value="1">
                    <em>Partidas</em>
                </MenuItem>
                <MenuItem value="2">
                    <em>Estatísticas</em>
                </MenuItem>
            </Select>
        </form>
    )
}
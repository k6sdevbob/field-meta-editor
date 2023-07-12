import NotificationWrapper from "./notify";
import { useCombineProviders } from "./store/utils";
import { useMainStoreProvider } from "./store/mainStore";


const App = () => {
    const ViewStoreProvider = useMainStoreProvider();
    const CombinedProvider = useCombineProviders(ViewStoreProvider);

    return (
        <CombinedProvider>
            <NotificationWrapper>
                {/* element */}
                <></>
            </NotificationWrapper>
        </CombinedProvider>
    );
};


export default App;
